package visitor

import (
	"encoding/json"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	parser "github.com/kulics-works/feel-go/parser/generate"
)

const (
	Wrap      = "\r\n"
	Terminate = ";"

	Byte = "byte"
	Any  = "interface{}"
	Int  = "int"
	Num  = "float64"
	I8   = "int8"
	I16  = "int16"
	I32  = "int32"
	I64  = "int64"

	U8  = "uint8"
	U16 = "uint16"
	U32 = "uint32"
	U64 = "uint64"

	F32 = "float32"
	F64 = "float64"

	Bool = "bool"
	T    = "true"
	F    = "false"

	Chr = "rune"
	Str = "string"
	Lst = "List"
	Set = "Set"
	Dic = "Dictionary"

	BlockLeft  = "{"
	BlockRight = "}"

	Func  = "func "
	Var   = "var "
	Const = "const "
	Type  = "type "
	Chan  = "chan "
)

type any = interface{}
type str = string

type errorListener struct {
	*antlr.DefaultErrorListener

	file string
	Err  error
}

func NewErrorListener(file string) *errorListener {
	return &errorListener{
		file: file,
	}
}

func (me *errorListener) Error() string {
	if me.Err == nil {
		return ""
	}
	return me.Err.Error()
}

func (me *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol any, line, column int, msg string, e antlr.RecognitionException) {
	me.Err = fmt.Errorf("[ERR %s:%d:%d] %s", me.file, line, column, msg)
}

type KVisitor struct {
	parser.BaseFeelParserVisitor

	AllIDSet     *set_str
	CurrentIDSet *stack_str

	self Parameter
}

func NewKVisitor() KVisitor {
	set := new_set_str()
	stack := new_stack_str()
	stack.push(set)
	return KVisitor{AllIDSet: set, CurrentIDSet: stack}
}

type Result struct {
	Data       any
	Text       string
	Permission string
	IsVirtual  bool
	IsDefine   bool
}
type set_str struct {
	set map[str]struct{}
}

func new_set_str() *set_str { return &set_str{make(map[str]struct{})} }
func (me *set_str) contains(id str) bool {
	_, has := me.set[id]
	return has
}
func (me *set_str) add(id str) {
	me.set[id] = struct{}{}
}
func (me *set_str) except_with(target *set_str) {
	for k := range target.set {
		if me.contains(k) {
			delete(me.set, k)
		}
	}
}

type stack_str struct {
	stack []*set_str
}

func new_stack_str() *stack_str { return &stack_str{make([]*set_str, 0)} }

func (me *stack_str) peek() *set_str {
	return me.stack[len(me.stack)-1]
}
func (me *stack_str) push(new *set_str) {
	me.stack = append(me.stack, new)
}
func (me *stack_str) pop() {
	me.stack = me.stack[:len(me.stack)-1]
}

func (me *KVisitor) has_id(id str) bool {
	return me.AllIDSet.contains(id) || me.CurrentIDSet.peek().contains(id)
}
func (me *KVisitor) add_id(id str) {
	me.CurrentIDSet.peek().add(id)
}
func (me *KVisitor) add_current_set() {
	for k := range me.CurrentIDSet.peek().set {
		me.AllIDSet.add(k)
	}
	me.CurrentIDSet.push(new_set_str())
}
func (me *KVisitor) delete_current_set() {
	// printJSON(me.AllIDSet)
	me.AllIDSet.except_with(me.CurrentIDSet.peek())
	// printJSON(me.AllIDSet)
	me.CurrentIDSet.pop()
}

func printJSON(i any) {
	by, _ := json.Marshal(i)
	fmt.Println(string(by))
}

func (me *KVisitor) Visit(tree antlr.ParseTree) any {
	return tree.Accept(me)
}

func (me *KVisitor) VisitChildren(tree antlr.RuleNode) any {
	return tree.Accept(me)
}

// func (me *KVisitor) VisitTerminal(tree antlr.TerminalNode) any {
// 	return tree.Accept(me)
// }

// func (me *KVisitor) VisitErrorNode(tree antlr.ErrorNode) any {
// 	return tree.Accept(me)
// }

func (me *KVisitor) VisitProgram(ctx *parser.ProgramContext) any {
	obj := ""
	for _, item := range ctx.AllStatement() {
		obj += me.Visit(item).(string)
	}
	return obj
}

func (me *KVisitor) VisitId(ctx *parser.IdContext) any {
	r := Result{Data: "var"}
	first := me.Visit(ctx.GetChild(0).(antlr.ParseTree)).(Result)
	r.Permission = first.Permission
	r.Text = first.Text
	r.IsVirtual = first.IsVirtual
	if ctx.GetChildCount() >= 2 {
		for i := 1; i < ctx.GetChildCount(); i++ {
			other := me.Visit(ctx.GetChild(i).(antlr.ParseTree)).(Result)
			r.Text += "_" + other.Text
		}
	}
	// todo
	// if keywords.Exists({t -> t == r.Text}) {
	// 	r.Text = "@" + r.Text
	// }
	return r
}

func (me *KVisitor) VisitIdItem(ctx *parser.IdItemContext) any {
	r := Result{Data: "var"}
	if ctx.TypeBasic() != nil {
		r.Permission = "public"
		r.Text += ctx.TypeBasic().GetText()
	} else if ctx.TypeAny() != nil {
		r.Permission = "public"
		r.Text += ctx.TypeAny().GetText()
	} else if ctx.GetOp().GetTokenType() == parser.FeelLexerIDPublic {
		r.Permission = "public"
		r.Text += ctx.GetOp().GetText()
		// r.IsVirtual = r.Text[0].is Upper()
	} else if ctx.GetOp().GetTokenType() == parser.FeelLexerIDPrivate {
		r.Permission = "protected"
		r.Text += ctx.GetOp().GetText()
		// r.IsVirtual = r.Text[r.Text.find first({it -> it >< '_'})].is Upper()
	}
	return r
}

func (me *KVisitor) VisitVarId(ctx *parser.VarIdContext) any {
	if ctx.Discard() != nil {
		return "_"
	} else {
		var id = me.Visit(ctx.Id()).(Result).Text
		if me.has_id(id) {
			// r.isDefine = true
		} else {
			me.add_id(id)
		}
		return id
		// if ctx.TypeType() != nil {
		// 	return me.Visit(ctx.TypeType()) + " " + id
		// } else {
		// 	return "var " + id
		// }
	}
}

func (me *KVisitor) VisitConstId(ctx *parser.ConstIdContext) any {
	if ctx.Discard() != nil {
		return "_"
	} else {
		var id = me.Visit(ctx.Id()).(Result).Text
		if me.has_id(id) {
			// r.isDefine = true
		} else {
			me.add_id(id)
		}
		return id
		// if ctx.TypeType() != nil {
		// 	return me.Visit(ctx.TypeType()) + " " + id
		// } else {
		// 	return "var " + id
		// }
	}
}

// func (me *KVisitor) VisitIdExpression(ctx *parser.IdExpressionContext) any {
// 	r := Result{Data: "var"}
// 	if len(ctx.AllIdExprItem()) > 1 {
// 		r.Text = ""
// 		for i, v := range ctx.AllIdExprItem() {
// 			subID := me.Visit(v).(Result).Text
// 			if i != 0 {
// 				r.Text += ", " + subID
// 			} else {
// 				r.Text += subID
// 			}
// 			if me.has_id(subID) {
// 				r.IsDefine = true
// 			} else {
// 				me.add_id(subID)
// 			}
// 		}
// 		r.Text += ""
// 	} else {
// 		r = me.Visit(ctx.IdExprItem(0)).(Result)
// 		if me.has_id(r.Text) {
// 			r.IsDefine = true
// 		} else {
// 			me.add_id(r.Text)
// 		}
// 	}
// 	return r
// }

// func (me *KVisitor) VisitIdExprItem(ctx *parser.IdExprItemContext) any {
// 	return me.Visit(ctx.GetChild(0).(antlr.ParseTree))
// }
