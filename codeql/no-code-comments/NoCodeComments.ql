// SPDX-FileCopyrightText: 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

/**
 * @name Code comment
 * @description Code has no comments. A comment needs a legal marker, a functional directive, an https link that proves it, the Go Authors license header of a forked file, a doc comment of the public API, or the 1 doc comment of a library package that staticcheck ST1000 requires.
 * @kind problem
 * @problem.severity error
 * @id go/no-code-comments
 */

import go

predicate proven(Comment comment) {
  exists(CommentGroup group, int index, int later |
    comment = group.getComment(index) and
    later >= index and
    group.getComment(later).getText().regexpMatch("(?s).*(https://|SPDX-|REUSE-Ignore).*")
  )
}

predicate directive(Comment comment) {
  comment.getText().regexpMatch("(?s)(go:[a-z]+|\\+build|nolint|line )\\b.*|\\s*(Deprecated: .*|(//\\s*)?indirect\\s*)")
}

bindingset[name]
predicate exportedName(string name) { name.regexpMatch("[A-Z].*") }

predicate exported(Documentable element) {
  exportedName(element.(FuncDecl).getName())
  or
  exportedName(element.(TypeSpec).getName())
  or
  exportedName(element.(ValueSpec).getAName())
  or
  exported(element.(GenDecl).getASpec())
  or
  exportedName(element.(FieldDecl).getNameExpr(_).(Ident).getName())
  or
  element instanceof EmbeddedFieldDecl
  or
  exportedName(element.(MethodSpec).getNameExpr().(Ident).getName())
  or
  element instanceof EmbeddingSpec
}

predicate documentsPublicApi(Comment comment) {
  not comment.getFile().getRelativePath().regexpMatch("((internal|cmd|scripts|examples)/.*|.*_test\\.go)") and
  exists(DocComment doc | comment = doc.getAComment() and exported(doc.getDocumentedElement()))
}

predicate documentsPackage(GoFile file) {
  exists(file.getDocumentation()) and
  not file.getPackageName() = "main" and
  not file.getBaseName().matches("%\\_test.go")
}

predicate packageDoc(Comment comment) {
  exists(GoFile file |
    documentsPackage(file) and
    comment = file.getDocumentation().getComment(_) and
    not exists(GoFile earlier |
      documentsPackage(earlier) and
      earlier.getParentContainer() = file.getParentContainer() and
      earlier.getPackageName() = file.getPackageName() and
      earlier.getRelativePath() < file.getRelativePath()
    )
  )
}

predicate goAuthorsLicense(Comment comment) {
  exists(CommentGroup header |
    comment = header.getAComment() and
    header.getAComment().getText().regexpMatch("\\s*Copyright \\d{4} The Go Authors\\. All rights reserved\\.\\s*")
  )
}

from Comment comment
where
  exists(comment.getFile().getRelativePath()) and
  not goAuthorsLicense(comment) and
  not proven(comment) and
  not directive(comment) and
  not documentsPublicApi(comment) and
  not packageDoc(comment)
select comment, "Code has no comments. Put the meaning in names, types and tests."
