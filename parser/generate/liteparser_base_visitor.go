// Code generated from LiteParser.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // LiteParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLiteParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLiteParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitExportStatement(ctx *ExportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitNamespaceSupportStatement(ctx *NamespaceSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeAliasStatement(ctx *TypeAliasStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeRedefineStatement(ctx *TypeRedefineStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitEnumStatement(ctx *EnumStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitEnumSupportStatement(ctx *EnumSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitNamespaceVariableStatement(ctx *NamespaceVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitNamespaceConstantStatement(ctx *NamespaceConstantStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitNamespaceFunctionStatement(ctx *NamespaceFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageStatement(ctx *PackageStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageFieldStatement(ctx *PackageFieldStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageSupportStatement(ctx *PackageSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitIncludeStatement(ctx *IncludeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageNewStatement(ctx *PackageNewStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageVariableStatement(ctx *PackageVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageFunctionStatement(ctx *PackageFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageControlSubStatement(ctx *PackageControlSubStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageEventStatement(ctx *PackageEventStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPackageImplementStatement(ctx *PackageImplementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitImplementStatement(ctx *ImplementStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitImplementSupportStatement(ctx *ImplementSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitImplementVariableStatement(ctx *ImplementVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitImplementFunctionStatement(ctx *ImplementFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitOverrideVariableStatement(ctx *OverrideVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitOverrideFunctionStatement(ctx *OverrideFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitProtocolStatement(ctx *ProtocolStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitProtocolSupportStatement(ctx *ProtocolSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitProtocolVariableStatement(ctx *ProtocolVariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitProtocolFunctionStatement(ctx *ProtocolFunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitFunctionStatement(ctx *FunctionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitReturnAwaitStatement(ctx *ReturnAwaitStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitYieldReturnStatement(ctx *YieldReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitYieldBreakStatement(ctx *YieldBreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitParameterClauseIn(ctx *ParameterClauseInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitParameterClauseOut(ctx *ParameterClauseOutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitParameter(ctx *ParameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitFunctionSupportStatement(ctx *FunctionSupportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeCaseStatement(ctx *JudgeCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCaseStatement(ctx *CaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCaseExprStatement(ctx *CaseExprStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeStatement(ctx *JudgeStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeElseStatement(ctx *JudgeElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeIfStatement(ctx *JudgeIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeElseIfStatement(ctx *JudgeElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopEachStatement(ctx *LoopEachStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopCaseStatement(ctx *LoopCaseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopElseStatement(ctx *LoopElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopJumpStatement(ctx *LoopJumpStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopContinueStatement(ctx *LoopContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCheckStatement(ctx *CheckStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitUsingStatement(ctx *UsingStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCheckErrorStatement(ctx *CheckErrorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCheckFinallyStatment(ctx *CheckFinallyStatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCheckReportStatement(ctx *CheckReportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitIteratorStatement(ctx *IteratorStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitVariableStatement(ctx *VariableStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitVariableDeclaredStatement(ctx *VariableDeclaredStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitChannelAssignStatement(ctx *ChannelAssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitAssignStatement(ctx *AssignStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitExpressionStatement(ctx *ExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitIdExpression(ctx *IdExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitIdExprItem(ctx *IdExprItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTupleExpression(ctx *TupleExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallExpression(ctx *CallExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTuple(ctx *TupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitExpressionList(ctx *ExpressionListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitAnnotationSupport(ctx *AnnotationSupportContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitAnnotation(ctx *AnnotationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitAnnotationList(ctx *AnnotationListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitAnnotationItem(ctx *AnnotationItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallFunc(ctx *CallFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallChannel(ctx *CallChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallElement(ctx *CallElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallPkg(ctx *CallPkgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallNew(ctx *CallNewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitOrElse(ctx *OrElseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeConversion(ctx *TypeConversionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPkgAssign(ctx *PkgAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPkgAssignElement(ctx *PkgAssignElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitListAssign(ctx *ListAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitSetAssign(ctx *SetAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitDictionaryAssign(ctx *DictionaryAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallAwait(ctx *CallAwaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCallAsync(ctx *CallAsyncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitList(ctx *ListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitSet(ctx *SetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitDictionary(ctx *DictionaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitDictionaryElement(ctx *DictionaryElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitSlice(ctx *SliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitSliceFull(ctx *SliceFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitSliceStart(ctx *SliceStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitSliceEnd(ctx *SliceEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitNameSpaceItem(ctx *NameSpaceItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTemplateDefine(ctx *TemplateDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTemplateDefineItem(ctx *TemplateDefineItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTemplateCall(ctx *TemplateCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLambda(ctx *LambdaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLambdaIn(ctx *LambdaInContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPkgAnonymous(ctx *PkgAnonymousContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPkgAnonymousAssign(ctx *PkgAnonymousAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPkgAnonymousAssignElement(ctx *PkgAnonymousAssignElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitFunctionExpression(ctx *FunctionExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPlusMinus(ctx *PlusMinusContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitNegate(ctx *NegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwiseNotExpression(ctx *BitwiseNotExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLinq(ctx *LinqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLinqHeadItem(ctx *LinqHeadItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLinqItem(ctx *LinqItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitStringExpression(ctx *StringExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitStringExpressionElement(ctx *StringExpressionElementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeExpression(ctx *JudgeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeExpressionElseStatement(ctx *JudgeExpressionElseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeExpressionIfStatement(ctx *JudgeExpressionIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeExpressionElseIfStatement(ctx *JudgeExpressionElseIfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeCaseExpression(ctx *JudgeCaseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCaseExpressionStatement(ctx *CaseExpressionStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopExpression(ctx *LoopExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopEachExpression(ctx *LoopEachExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLoopElseExpression(ctx *LoopElseExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCheckExpression(ctx *CheckExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCheckErrorExpression(ctx *CheckErrorExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitDataStatement(ctx *DataStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitFloatExpr(ctx *FloatExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitIntegerExpr(ctx *IntegerExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeNotNull(ctx *TypeNotNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeType(ctx *TypeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeReference(ctx *TypeReferenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeNullable(ctx *TypeNullableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeArray(ctx *TypeArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeList(ctx *TypeListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeSet(ctx *TypeSetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeDictionary(ctx *TypeDictionaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeStack(ctx *TypeStackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeQueue(ctx *TypeQueueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeChannel(ctx *TypeChannelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypePackage(ctx *TypePackageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeFunction(ctx *TypeFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeAny(ctx *TypeAnyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeFunctionParameterClause(ctx *TypeFunctionParameterClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitTypeBasic(ctx *TypeBasicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitNilExpr(ctx *NilExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudgeType(ctx *JudgeTypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwise(ctx *BitwiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwiseAnd(ctx *BitwiseAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwiseOr(ctx *BitwiseOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwiseNot(ctx *BitwiseNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwiseXor(ctx *BitwiseXorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwiseLeftShift(ctx *BitwiseLeftShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitBitwiseRightShift(ctx *BitwiseRightShiftContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitJudge(ctx *JudgeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitAssign(ctx *AssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitAdd(ctx *AddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitMul(ctx *MulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitPow(ctx *PowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitCall(ctx *CallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitWave(ctx *WaveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitIdItem(ctx *IdItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitEnd(ctx *EndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitMore(ctx *MoreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLeft_brace(ctx *Left_braceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitRight_brace(ctx *Right_braceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLeft_paren(ctx *Left_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitRight_paren(ctx *Right_parenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitLeft_brack(ctx *Left_brackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLiteParserVisitor) VisitRight_brack(ctx *Right_brackContext) interface{} {
	return v.VisitChildren(ctx)
}
