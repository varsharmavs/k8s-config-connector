// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mockspannerinstance

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/common/projects"
)

type spannerInstanceName struct {
	Project      *projects.ProjectData
	InstanceName string
}

func (n *spannerInstanceName) String() string {
	return "projects/" + n.Project.ID + "/instances/" + n.InstanceName
}

func (s *MockService) parseInstanceName(name string) (*spannerInstanceName, error) {
	tokens := strings.Split(name, "/")

	if len(tokens) == 4 && tokens[0] == "projects" && tokens[2] == "instances" {
		project, err := s.Projects.GetProjectByID(tokens[1])
		if err != nil {
			return nil, err
		}

		name := &spannerInstanceName{
			Project:      project,
			InstanceName: tokens[3],
		}

		return name, nil
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "name %q is not valid", name)
	}
}
