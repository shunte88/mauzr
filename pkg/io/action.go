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

package io

import "time"

// Action represents a defined action that can be executed later
type Action func() error

// Execute the given Actions. Remaining Actions are not executed if an Action
// fails. After Actions are handled, all cleanup actions all executed.
func Execute(actions, cleanup []Action) error {
	var firstError error
	for _, action := range actions {
		if err := action(); err != nil {
			firstError = err
			break
		}
	}

	for _, action := range cleanup {
		if err := action(); err != nil && firstError == nil {
			firstError = err
		}
	}
	return firstError
}

// Sleep create an Action that sleep for a given time.
func Sleep(duration time.Duration) Action {
	return func() error {
		time.Sleep(duration)
		return nil
	}
}

// NoOperation is an action that does nothing.
func NoOperation() error {
	return nil
}
