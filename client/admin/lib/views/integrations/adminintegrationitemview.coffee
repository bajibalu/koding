kd             = require 'kd'
JView          = require 'app/jview'
KDButtonView   = kd.ButtonView
KDListItemView = kd.ListItemView


module.exports = class AdminIntegrationItemView extends KDListItemView

  JView.mixin @prototype

  constructor: (options = {}, data) ->

    options.cssClass = 'integration-view'

    super options, data

    @addButton = new KDButtonView
      cssClass : 'solid compact green add'
      title    : 'Add'


  pistachio: ->
    { name, desc, logo } = @getData()

    return """
      <img src="#{logo}" />
      <p>#{name}</p>
      <span>#{desc}</span>
      {{> @addButton}}
    """
