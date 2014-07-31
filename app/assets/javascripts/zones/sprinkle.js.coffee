class Status
  initialize: =>
    @observer = new MutationObserver(@observed)
    @observer.observe(@element(), {
      attributes: true
      })

    @remaining = document.createElement('div')
    @remaining.innerHTML = """
    <time as='Sprinkles.Moment'></time>
    """

    dt = @element().getAttribute('datetime')
    if dt
      @opened(dt)

  addEventListeners: =>
    @element().querySelector('select').addEventListener "change", @submit

  removeEventListeners: =>
    @observer.disconnect()
    @element().querySelector('select').removeEventListener "change", @submit


  observed: (mutations) =>
    mutations.forEach (m) =>
      dt = @element().getAttribute(m.attributeName)
      if dt?
        @opened(dt)
      else
        @closed(dt)

  opened: (dt) =>
    @remaining.querySelector('time').setAttribute('datetime', dt)
    @form = @element().querySelector('form.sprinkle')
    @form.parentElement.replaceChild(@remaining, @form)

  closed: (dt) =>
    @remaining.parentElement.replaceChild(@form, @remaining)

  submit: (e) =>
    fd = new FormData(e.target.form)
    fd.append(
      document.querySelector("[name=csrf-param]").getAttribute('content'),
      document.querySelector("[name=csrf-token]").getAttribute('content')
      )
    xhr = new XHR(e.target)
    xhr.send(e.target.form.action, "POST", fd)
    e.target.children[0].setAttribute('selected', true)

Shiny.Models.add Status, "Sprinkles.Status"
