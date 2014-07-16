instance = undefined

class Watcher
  constructor: (target, config = {}) ->
    @observer = new MutationObserver(@observed)
    @observer.observe(target, config)

  observed: (mutations) =>
    mutations.forEach (mutation) =>
      if mutation.type == 'attributes'
        Shiny.God.update(target)
      else
        @add(mutation.addedNodes)
        @destroy(mutation.removedNodes)


  add: (nodes) =>
    for node in nodes
      continue unless Shiny.isDOM(node)
      if node.hasAttribute(Shiny.attributeName)
        Shiny.God.create(node, node.getAttribute(Shiny.attributeName))

      for child in node.querySelectorAll("[#{Shiny.attributeName}]")
        Shiny.God.create(child, child.getAttribute(Shiny.attributeName))

  destroy: (nodes) =>
    for node in nodes
      continue unless Shiny.isDOM(node)
      if node.hasAttribute(Shiny.attributeName)
        Shiny.God.destroy(node)

      for child in node.querySelectorAll("[#{Shiny.attributeName}]")
        Shiny.God.destroy(child)

  inspect: (node) ->
    if Shiny.isDOM(node)
      found = node.querySelectorAll("[#{Shiny.attributeName}]")
      Shiny.God.create(el) for el in found

# !! **************************************** !! #

Shiny.Watcher = ->
  unless instance?
    i = 0
    target = null
    target = if Shiny.isDOM(arguments[i]) then arguments[i++] else document
    instance = new Watcher(target, arguments[i])

  instance

