@Shiny = {
  attributeName: 'as'
}

@Shiny.isDOM = (el) ->
  el instanceof HTMLDocument ||
  el instanceof HTMLElement

listen = (e) ->
  if e.type && e.type == 'DOMContentLoaded'
    document.removeEventListener('DOMContentLoaded', listen)

  Shiny.Watcher(document, {
      attributes: true,
      subtree: true,
      childList: true,
      attributeFilter: [Shiny.attributeName],
      characterData: true
  })

  Shiny.Watcher().inspect(document.body)

if document.readyState == 'complete'
  listen()
else
  document.addEventListener('DOMContentLoaded', listen)


