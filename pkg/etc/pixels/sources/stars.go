/*
Copyright 2019 Alexander Sowitzki.

GNU Affero General Public License version 3 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/AGPL-3.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sources

import (
	"math/rand"

	"go.eqrx.net/mauzr/pkg/etc/pixels/color"
)

type stars struct {
	loopCommon
	lower, upper     color.RGBW
	factors, changes []float64
}

const (
	speed = 1 / 24
)

// NewStarts returns a Loop that lets the pixels imitate starts.
func NewStars(theme color.RGBW) Loop {
	return &stars{loopCommon{}, color.Off.MixWith(0.1, theme), theme, nil, nil}
}

// SetLength of the target pixel strip. May be called only once.
func (c *stars) SetLength(length int) {
	c.length = length
	c.factors = make([]float64, length)
	c.changes = make([]float64, length)
	for i := range c.factors {
		c.factors[i] = 1.0
		c.changes[i] = rand.Float64()*speed + 0.01
	}
}

// Peer the next generated color (Next invocation will return the same color).
func (c *stars) Peek() []color.RGBW {
	new := make([]color.RGBW, c.length)
	for i := range c.factors {
		new[i] = c.lower.MixWith(c.factors[i], c.upper)
	}
	return new
}

// Pop the next generated color (Next invocation will return the next color).
func (c *stars) Pop() []color.RGBW {
	new := c.Peek()
	for i := range c.factors {
		c.factors[i] += c.changes[i]
		switch {
		case c.factors[i] >= 1.0:
			c.factors[i] = 1.0
			c.changes[i] = -(rand.Float64()*speed + 0.01)
		case c.factors[i] <= 0.0:
			c.factors[i] = 0.0
			c.changes[i] = rand.Float64()*speed + 0.01
		}
	}
	return new
}