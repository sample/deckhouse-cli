// Copyright 2024 Flant JSC
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

package pull

import (
	"github.com/spf13/pflag"
)

func addFlags(flagSet *pflag.FlagSet) {
	flagSet.Int64VarP(
		&ImagesBundleChunkSizeGB,
		"images-bundle-chunk-size",
		"c",
		0,
		"Split resulting bundle file into chunks of 'size' GB",
	)
	flagSet.StringVar(
		&SourceRegistryRepo,
		"source",
		enterpriseEditionRepo,
		"Pull Deckhouse images from source registry. This is the default mode of operation.",
	)
	flagSet.StringVar(
		&SourceRegistryLogin,
		"source-login",
		"",
		"Source registry login.",
	)
	flagSet.StringVar(
		&SourceRegistryPassword,
		"source-password",
		"",
		"Source registry password.",
	)
	flagSet.StringVarP(
		&DeckhouseLicenseToken,
		"license",
		"l",
		"",
		"Pull Deckhouse images to local machine using license key. Shortcut for --source-login=license-token --source-password=<>.",
	)
	flagSet.StringVarP(
		&minVersionString,
		"min-version",
		"m",
		"",
		"Minimal Deckhouse release to copy. Cannot be above current Rock Solid release.",
	)
	flagSet.StringVar(
		&specificReleaseString,
		"release",
		"",
		"Specific Deckhouse release to copy. Conflicts with --min-version.",
	)
	flagSet.BoolVar(
		&DoGOSTDigest,
		"gost-digest",
		false,
		"Calculate GOST R 34.11-2012 STREEBOG digest for downloaded bundle",
	)
	flagSet.BoolVar(
		&DontContinuePartialPull,
		"no-pull-resume",
		false,
		"Do not continue last unfinished pull operation.",
	)
	flagSet.BoolVar(
		&NoModules,
		"no-modules",
		false,
		"Do not pull Deckhouse modules into bundle.",
	)
	flagSet.BoolVar(
		&TLSSkipVerify,
		"tls-skip-verify",
		false,
		"Disable TLS certificate validation.",
	)
	flagSet.BoolVar(
		&Insecure,
		"insecure",
		false,
		"Interact with registries over HTTP.",
	)
}