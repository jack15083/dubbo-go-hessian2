// Copyright 2016-2019 summerbuger@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package java_exception

type ObjectStreamException struct {
	SerialVersionUID     int64
	DetailMessage        string
	StackTrace           []StackTraceElement
	SuppressedExceptions []Throwabler
	Cause                Throwabler
}

func NewObjectStreamException(detailMessage string) *ObjectStreamException {
	return &ObjectStreamException{DetailMessage: detailMessage, StackTrace: []StackTraceElement{}}
}

func (e ObjectStreamException) Error() string {
	return e.DetailMessage
}

func (ObjectStreamException) JavaClassName() string {
	return "java.io.ObjectStreamException"
}
