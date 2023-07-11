// Copyright 2023 Google LLC. All Rights Reserved.
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
// GENERATED BY gen_go_data.go
// gen_go_data -package beta -var YAML_key blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/recaptchaenterprise/beta/key.yaml

package beta

// blaze-out/k8-fastbuild/genfiles/cloud/graphite/mmv2/services/google/recaptchaenterprise/beta/key.yaml
var YAML_key = []byte("info:\n  title: RecaptchaEnterprise/Key\n  description: The RecaptchaEnterprise Key resource\n  x-dcl-struct-name: Key\n  x-dcl-has-iam: false\npaths:\n  get:\n    description: The function used to get information about a Key\n    parameters:\n    - name: key\n      required: true\n      description: A full instance of a Key\n  apply:\n    description: The function used to apply information about a Key\n    parameters:\n    - name: key\n      required: true\n      description: A full instance of a Key\n  delete:\n    description: The function used to delete a Key\n    parameters:\n    - name: key\n      required: true\n      description: A full instance of a Key\n  deleteAll:\n    description: The function used to delete all Key\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\n  list:\n    description: The function used to list information about many Key\n    parameters:\n    - name: project\n      required: true\n      schema:\n        type: string\ncomponents:\n  schemas:\n    Key:\n      title: Key\n      x-dcl-id: projects/{{project}}/keys/{{name}}\n      x-dcl-parent-container: project\n      x-dcl-labels: labels\n      x-dcl-has-create: true\n      x-dcl-has-iam: false\n      x-dcl-read-timeout: 0\n      x-dcl-apply-timeout: 0\n      x-dcl-delete-timeout: 0\n      type: object\n      required:\n      - displayName\n      - project\n      properties:\n        androidSettings:\n          type: object\n          x-dcl-go-name: AndroidSettings\n          x-dcl-go-type: KeyAndroidSettings\n          description: Settings for keys that can be used by Android apps.\n          x-dcl-conflicts:\n          - webSettings\n          - iosSettings\n          properties:\n            allowAllPackageNames:\n              type: boolean\n              x-dcl-go-name: AllowAllPackageNames\n              description: If set to true, it means allowed_package_names will not\n                be enforced.\n            allowedPackageNames:\n              type: array\n              x-dcl-go-name: AllowedPackageNames\n              description: 'Android package names of apps allowed to use the key.\n                Example: ''com.companyname.appname'''\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        createTime:\n          type: string\n          format: date-time\n          x-dcl-go-name: CreateTime\n          readOnly: true\n          description: The timestamp corresponding to the creation of this Key.\n          x-kubernetes-immutable: true\n        displayName:\n          type: string\n          x-dcl-go-name: DisplayName\n          description: Human-readable display name of this key. Modifiable by user.\n        iosSettings:\n          type: object\n          x-dcl-go-name: IosSettings\n          x-dcl-go-type: KeyIosSettings\n          description: Settings for keys that can be used by iOS apps.\n          x-dcl-conflicts:\n          - webSettings\n          - androidSettings\n          properties:\n            allowAllBundleIds:\n              type: boolean\n              x-dcl-go-name: AllowAllBundleIds\n              description: If set to true, it means allowed_bundle_ids will not be\n                enforced.\n            allowedBundleIds:\n              type: array\n              x-dcl-go-name: AllowedBundleIds\n              description: 'iOS bundle ids of apps allowed to use the key. Example:\n                ''com.companyname.productname.appname'''\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n        labels:\n          type: object\n          additionalProperties:\n            type: string\n          x-dcl-go-name: Labels\n          description: See [Creating and managing labels](https://cloud.google.com/recaptcha-enterprise/docs/labels).\n        name:\n          type: string\n          x-dcl-go-name: Name\n          description: The resource name for the Key in the format \"projects/{project}/keys/{key}\".\n          x-kubernetes-immutable: true\n          x-dcl-server-generated-parameter: true\n        project:\n          type: string\n          x-dcl-go-name: Project\n          description: The project for the resource\n          x-kubernetes-immutable: true\n          x-dcl-references:\n          - resource: Cloudresourcemanager/Project\n            field: name\n            parent: true\n        testingOptions:\n          type: object\n          x-dcl-go-name: TestingOptions\n          x-dcl-go-type: KeyTestingOptions\n          description: Options for user acceptance testing.\n          x-kubernetes-immutable: true\n          properties:\n            testingChallenge:\n              type: string\n              x-dcl-go-name: TestingChallenge\n              x-dcl-go-type: KeyTestingOptionsTestingChallengeEnum\n              description: 'For challenge-based keys only (CHECKBOX, INVISIBLE), all\n                challenge requests for this site will return nocaptcha if NOCAPTCHA,\n                or an unsolvable challenge if UNSOLVABLE_CHALLENGE. Possible values:\n                TESTING_CHALLENGE_UNSPECIFIED, NOCAPTCHA, UNSOLVABLE_CHALLENGE'\n              x-kubernetes-immutable: true\n              x-dcl-server-default: true\n              enum:\n              - TESTING_CHALLENGE_UNSPECIFIED\n              - NOCAPTCHA\n              - UNSOLVABLE_CHALLENGE\n            testingScore:\n              type: number\n              format: double\n              x-dcl-go-name: TestingScore\n              description: All assessments for this Key will return this score. Must\n                be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive.\n              x-kubernetes-immutable: true\n        webSettings:\n          type: object\n          x-dcl-go-name: WebSettings\n          x-dcl-go-type: KeyWebSettings\n          description: Settings for keys that can be used by websites.\n          x-dcl-conflicts:\n          - androidSettings\n          - iosSettings\n          required:\n          - integrationType\n          properties:\n            allowAllDomains:\n              type: boolean\n              x-dcl-go-name: AllowAllDomains\n              description: If set to true, it means allowed_domains will not be enforced.\n            allowAmpTraffic:\n              type: boolean\n              x-dcl-go-name: AllowAmpTraffic\n              description: If set to true, the key can be used on AMP (Accelerated\n                Mobile Pages) websites. This is supported only for the SCORE integration\n                type.\n            allowedDomains:\n              type: array\n              x-dcl-go-name: AllowedDomains\n              description: 'Domains or subdomains of websites allowed to use the key.\n                All subdomains of an allowed domain are automatically allowed. A valid\n                domain requires a host and must not include any path, port, query\n                or fragment. Examples: ''example.com'' or ''subdomain.example.com'''\n              x-dcl-send-empty: true\n              x-dcl-list-type: list\n              items:\n                type: string\n                x-dcl-go-type: string\n            challengeSecurityPreference:\n              type: string\n              x-dcl-go-name: ChallengeSecurityPreference\n              x-dcl-go-type: KeyWebSettingsChallengeSecurityPreferenceEnum\n              description: 'Settings for the frequency and difficulty at which this\n                key triggers captcha challenges. This should only be specified for\n                IntegrationTypes CHECKBOX and INVISIBLE. Possible values: CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED,\n                USABILITY, BALANCE, SECURITY'\n              x-dcl-server-default: true\n              enum:\n              - CHALLENGE_SECURITY_PREFERENCE_UNSPECIFIED\n              - USABILITY\n              - BALANCE\n              - SECURITY\n            integrationType:\n              type: string\n              x-dcl-go-name: IntegrationType\n              x-dcl-go-type: KeyWebSettingsIntegrationTypeEnum\n              description: 'Required. Describes how this key is integrated with the\n                website. Possible values: SCORE, CHECKBOX, INVISIBLE'\n              x-kubernetes-immutable: true\n              enum:\n              - SCORE\n              - CHECKBOX\n              - INVISIBLE\n")

// 8345 bytes
// MD5: 3eef1c4729975c1823cfb4be67af327a
