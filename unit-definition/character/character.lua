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
    choices.insert({
        name='name',
        type='string'
    })
    choices.insert({
      name='strength',
      type='int'
    })
    choices.insert({
        name='constitution',
        type='int'
    })
    choices.insert({
        name='dexterity',
        type='int'
    })
    choices.insert({
        name='intelligence',
        type='int'
    })
    choices.insert({
        name='charisma',
        type='int'
    })
    choices.insert({
        name='wisdom',
        type='int'
    })
end