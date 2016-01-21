// params_test.go
//
// Copyright (C) 2015 Selection Pressure LLC
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.


package ionconnect

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/ion-channel/ion-connect/Godeps/_workspace/src/github.com/codegangsta/cli"
  "flag"
)

var _ = Describe("Params", func() {
  var (
		context *cli.Context
    config Command
    set *flag.FlagSet
	)

  BeforeEach(func() {
    Debug = true
    set = flag.NewFlagSet("set", 0)
    command := cli.Command{Name: "scan-git"}
    context = cli.NewContext(nil, set, nil)
    context.Command = command

    config,_ = GetConfig().FindSubCommandConfig("scanner", "scan-git")
  })

  Context("When generating Post Params", func() {
    It("should populate the fields from the context flags", func() {
      args := []string{"ernie"}
      Expect(PostParams{}.Generate(args, config.Args)).To(Equal(PostParams{Project:"ernie"}))
    })
  })

  Context("When generating Get Params", func() {
    It("should populate the fields from the context flags", func() {
      args := []string{"ernie"}
      Expect(GetParams{}.Generate(args, config.Args)).To(Equal(GetParams{Project:"ernie"}))
    })
  })

  AfterEach(func() {
    Debug = false
  })
})
