// SPDX-FileCopyrightText: 2026 Xquik contributors
// SPDX-License-Identifier: Apache-2.0

/**
 * @name Code comment
 * @description Code has no comments. A comment needs a legal marker, a functional directive, an https link that proves it, or it documents the public API.
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
  or
  element instanceof File
}

predicate documentsPublicApi(Comment comment) {
  not comment.getFile().getRelativePath().regexpMatch("((internal|cmd|scripts|examples)/.*|.*_test\\.go)") and
  exists(DocComment doc | comment = doc.getAComment() and exported(doc.getDocumentedElement()))
}

predicate vendored(File file) {
  exists(Comment notice |
    notice.getFile() = file and
    notice.getText().regexpMatch("\\s*Copyright \\d{4} The Go Authors\\. All rights reserved\\.\\s*")
  )
}

from Comment comment
where
  exists(comment.getFile().getRelativePath()) and
  not vendored(comment.getFile()) and
  not proven(comment) and
  not directive(comment) and
  not documentsPublicApi(comment)
select comment, "Code has no comments. Put the meaning in names, types and tests."
