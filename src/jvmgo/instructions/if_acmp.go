package instructions

import "jvmgo/rtda"

// Branch if reference comparison succeeds 
type if_acmpeq struct {
    branch int16
}
func (self *if_acmpeq) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_acmpeq) execute(thread *rtda.Thread) {
    if _acmp(thread) {
        // todo
    }
}

type if_acmpne struct {
    branch int16
}
func (self *if_acmpne) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}
func (self *if_acmpne) execute(thread *rtda.Thread) {
    if !_acmp(thread) {
        // todo
    }
}

func _acmp(thread *rtda.Thread) (bool) {
    stack := thread.CurrentFrame().OperandStack()
    ref1 := stack.PopRef()
    ref2 := stack.PopRef()
    return ref1 == ref2 // todo
}
