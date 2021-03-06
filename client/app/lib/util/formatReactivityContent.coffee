transformTagTokens = require './transformTagTokens'
transformEmails = require './transformEmails'
formatQuotes = require './formatQuotes'
formatBlockquotes = require './formatBlockquotes'
applyMarkdown = require './applyMarkdown'
expandUsernames = require './expandReactivityUsernames'
markdownUrls = require './markdownUrls'
Encoder = require 'htmlencode'

module.exports = (body = '', markdownOptions = {}) ->

  fns = [
    transformTagTokens
    transformEmails
    formatQuotes
    formatBlockquotes
    markdownUrls
  ]

  body = fn body for fn in fns
  body = Encoder.htmlDecode body
  body = applyMarkdown body, markdownOptions
  body = expandUsernames body, 'code, a'

  return body
