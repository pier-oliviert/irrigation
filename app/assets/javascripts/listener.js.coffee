class Listener
  constructor: ->
    @socket = new WebSocket("ws://#{location.hostname}:21343")
    @socket.onmessage = @received
    @openHTML = document.createElement('div')
    @openHTML.innerHTML = """
    <div class='loading'>
    <time as='Sprinkles.Moment'></time>
    <a href='javascript:void(0)' class='cancel'>Annuler</a>
    """

  addEventListeners: =>

  removeEventListeners: =>

  received: (e) =>
    cmd = JSON.parse(e.data)
    if cmd.status == "open"
      @opened(cmd)
    else
      @closed(cmd)

  opened: (cmd) =>
    z = document.getElementById("zone_#{cmd.id}")
    t = @openHTML.querySelector('time')
    t.setAttribute('datetime', cmd.close_at)
    @formHTML = z.querySelector('form.sprinkle')
    @formHTML.parentElement.replaceChild(@openHTML, @formHTML)

  closed: (cmd) =>
    if @formHTML?
      @openHTML.parentElement.replaceChild(@formHTML, @openHTML)
      @formHTML = null



Shiny.Models.add Listener, "Listener"
