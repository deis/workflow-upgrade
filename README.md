# Workflow Upgrade
[![Docker Repository on Quay](https://quay.io/repository/deisci/workflow-upgrade/status "Docker Repository on Quay")](https://quay.io/repository/deisci/workflow-upgrade)

Deis (pronounced DAY-iss) Workflow is an open source Platform as a Service (PaaS) that adds a developer-friendly layer to any [Kubernetes](http://kubernetes.io) cluster, making it easy to deploy and manage applications on your own servers.

For more information about the Deis Workflow, please visit the main project page at https://github.com/deis/workflow.

We welcome your input! If you have feedback, please [submit an issue][issues]. If you'd like to participate in development, please read the "Development" section below and [submit a pull request][prs].

# About
The Workflow Upgrade service checks for the mismatch between the version of the daemonset pods and the [daemonset][daemons] and deletes the pods such that new pods come up with the latest changes. This is necessary until kubernetes has implemented the update strategy which is tracked through https://github.com/kubernetes/kubernetes/issues/22543.

Workflow Upgrade will be run as kubernetes job after a successful upgrade of the workflow to a new release version.

# License

Copyright 2013, 2014, 2015, 2016 Engine Yard, Inc.

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

[issues]: https://github.com/deis/workflow/issues
[prs]: https://github.com/deis/workflow/pulls
[daemons]: http://kubernetes.io/docs/admin/daemons/
