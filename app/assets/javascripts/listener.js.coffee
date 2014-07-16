class Listener
  constructor: ->
    @socket = new WebSocket('ws://localhost:21343')
    @socket.onmessage = @received

  addEventListeners: =>

  removeEventListeners: =>


  received: (e) ->
    console.log e

Shiny.Models.add Listener, "Listener"
