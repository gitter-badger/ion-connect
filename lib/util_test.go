package ionconnect

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/ion-channel/ion-connect/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
)

var _ = Describe("Util", func() {
  var (

  )

  BeforeEach(func() {
    Debug = true
  })

  Context("When the debug flag is set", func() {
    It("should write out debug statements", func() {
        Expect(func(){Debugln("testing")}).ShouldNot(Panic())
        Expect(func(){Debugf("testing %s", "f")}).ShouldNot(Panic())
    })
  })

  Context("When a checking for a file or folder", func() {
    It("return false if it doesn't exist", func() {
        Expect(PathExists("/aint/real")).To(BeFalse())
    })
    It("return true if it exists", func() {
        path, _ := tilde.Expand("~")
        Expect(PathExists(path)).To(BeTrue())
    })
  })

  AfterEach(func() {
    Debug = false
  })
})
