_G.core = {}

_G.core.filterBy = function(variantList, exceptionsDict)
    local resultList = {}
    for i, v in ipairs(variantList) do
        if exceptionsDict[v] ~= true then
            table.insert(resultList, v)
        end
    end
end

_G.core.merge = function(target, source)
  for k, v in pairs(source) do
    target[k] = v
  end
end

_G.core.findIndex = function(list, value)
    local index = 0
    for k, v in pairs(list) do
        if v == value then
            index = k
        end
    end
    return index
end

_G.core.getProp = function(ctx, field)
    return ctx.command.data[field]
end

_G.core.getVal = function(ctx, field)
    return ctx.command.context[field]
end

_G.core.getCommandContext = function(ctx)
    return ctx.command.context
end

function dispatch(ctx, commandName, data, context, isLocal, phase)
    local payload = {
        command=commandName,
        data={},
        phase=phase,
        context=ctx.command.context
    }
    if data then
        payload.data = data
    end
    if context then
        payload.context = _G.core.merge(payload.context, context)
    end
    if isLocal then
        payload.scope = ctx.scope
    end
    table.insert(ctx.next, payload)
end

_G.core.submitAll = function(ctx, commandName, data)
    dispatch(ctx, commandName, data, {}, false, 'submit')
end

_G.core.submitScoped = function(ctx, commandName, data, context)
    dispatch(ctx, commandName, data, context, true, 'submit')
end

_G.core.scheduleAll = function(ctx, commandName, data)
    dispatch(ctx, commandName, data, {}, false, 'schedule')
end

_G.core.scheduleScoped = function(ctx, commandName, data, context)
    dispatch(ctx, commandName, data, context, true, 'schedule')
end

_G.core.emitAll = function(ctx, commandName, data)
    dispatch(ctx, commandName, data, {}, false, 'emit')
end

_G.core.emitScoped = function(ctx, commandName, data, context)
    dispatch(ctx, commandName, data, context, true, 'emit')
end