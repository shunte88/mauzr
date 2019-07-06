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

package i2c

import "mauzr.eqrx.net/go/pkg/io"

// Device represents a device behind an I2C bus.
type Device interface {
	Open() io.Action
	Close() io.Action
	Write(source []byte) io.Action
	WriteRead(source []byte, destination []byte) io.Action
}

// NewDevice creates a new Device. this function can be overridden to mock the device.
var NewDevice = newNormalDevice
