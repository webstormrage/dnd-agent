onEmit = {}
onFulfill = {}
core = _G.core or {}
scope = 'unit-definition/abilities'

onEmit['Unit.addAbilities'] = function (ctx)
    local data = {
        optionName='strength',
        optionType='int',
    }
    local context = {
        unitId=core.getProp(ctx, 'unitId')
    }
    core.submitScoped(ctx, 'Option.scanf', data, context)
end

onFulfill['Option.scanf'] = function(ctx)
    local attr = core.getProp(ctx, 'optionName')
    local value = core.getProp(ctx, 'optionValue')
    local data = {
        unitId=core.getVal(ctx, 'unitId')
    }
    local context = {
        attr=attr,
        value=value
    }
    core.scheduleScoped(ctx, 'Unit.setAttributes', data, context)
end

onEmit['Unit.setAttributes'] = function(ctx)
    local attr = core.getVal(ctx, 'attr')
    local value = core.getVal(ctx, 'value')
    local attributes = core.getProp(ctx, 'attributes')
    attributes[attr] = math.min(value, 18)
    local data = {
        attributes=attributes
    }
    core.submitScoped(ctx, 'Unit.setAttributes', data)
end

onFulfill['Unit.setAttributes'] = function(ctx)
    local abilities = {'strength', 'constitution', 'dexterity', 'intelligence', 'charisma', 'wisdom'}
    local attr = core.getVal(ctx, 'attr')
    local index = core.findIndex(abilities, attr)
    if index == 0 then
        core.fulfillAll(ctx, 'Unit.addAbilities', {
            unitId=core.getVal(ctx, 'unitId')
        })
    else
        core.submitScoped(ctx, 'Option.scanf', {
            optionName='strength',
            optionType='int',
        })
    end
end