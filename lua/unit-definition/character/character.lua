function unitDefinition(attributes, inventory, options)
    attributes['name'] = options['name']
end

function optionsDefinition(attributes, choices)
    table.insert(choices, {
        name = 'name',
        type = 'string'
    })
    -- TODO добавить определение характера
end