class MomentImplementation
  addEventListeners: =>
    @update()

  removeEventListeners: =>

  update: =>
    dur = @duration()

    remaining = [
      @minutes(dur),
      @seconds(dur)
    ].filter(Boolean)

    if remaining.length > 0
      @element().innerText = remaining.join(", ")
      setTimeout(@update, 1000)
    else
      @element().innerText = "Fermeture en cours"

  duration: =>
    date = new Date(@element().getAttribute('datetime'))
    date - new Date()

  minutes: (dur) ->
    min = Math.floor(dur / (1000 * 60))
    if min > 1
      "#{min} minutes"
    else if min > 0
      "1 minute"

  seconds: (dur) ->
    s = Math.floor(dur / 1000) % 60
    if s > 1
      "#{s} secondes"
    else if s > 0
      "1 seconde"

Shiny.Models.add MomentImplementation, 'Sprinkles.Moment'
