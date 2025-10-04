function unitDefinition(attributes, equipment, options)
    attributes['strength'] = math.min(options['strength'], 18)
    attributes['constitution'] = math.min(options['constitution'], 18)
    attributes['dexterity'] = math.min(options['dexterity'], 18)
    attributes['intelligence'] = math.min(options['intelligence'], 18)
    attributes['charisma'] = math.min(options['charisma'], 18)
    attributes['wisdom'] = math.min(options['wisdom'], 18)
    attributes['name'] = options['name']
end

function optionsDefinition(attributes, choices)
    table.insert(choices, {
        name = 'name',
        type = 'string'
    })
    table.insert(choices, {
        name = 'strength',
        type = 'int'
    })
    table.insert(choices, {
        name = 'constitution',
        type = 'int'
    })
    table.insert(choices, {
        name = 'dexterity',
        type = 'int'
    })
    table.insert(choices, {
        name = 'intelligence',
        type = 'int'
    })
    table.insert(choices, {
        name = 'charisma',
        type = 'int'
    })
    table.insert(choices, {
        name = 'wisdom',
        type = 'int'
    })
end