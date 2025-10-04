function unitDefinition(attributes, inventory, options)
    -- Владения
    if attributes['proficiencies'] == nil then
        attributes['proficiencies'] = {}
    end
    attributes['proficiencies']['survival'] =  true
    attributes['proficiencies']['athletics'] = true
    attributes['proficiencies'][options['language']] = true
    attributes['proficiencies'][options['musical-skill']] = true

    -- Экипировка
    inventory['staff'] = (inventory['staff'] or 0) + 1
    inventory['hunting-trap'] = (inventory['hunting-trap'] or 0) + 1
    inventory['traveler-cloth'] = (inventory['traveler-cloth'] or 0) + 1
    inventory['coins'] = (inventory['coins'] or 0) + 1000
end

function optionsDefinition(attributes, choices)
    table.insert(choices, {
        name = 'language',
        type = 'select',
        options = {
            'language/abyssal',
            'language/celestial',
            'language/draconic',
            'language/deep-speech',
            'language/infernal',
            'language/primordial',
            'language/sylvan',
            'language/undercommon',
            'language/common',
            'language/dwarvish',
            'language/elvish',
            'language/giant',
            'language/gnomish',
            'language/goblin',
            'language/halfling',
            'language/orc'
        }
    })

    table.insert(choices, {
        name = 'musical-skill',
        type = 'select',
        options = {
            'bagpipes',
            'drum',
            'dulcimer',
            'flute',
            'lute',
            'lyre',
            'horn',
            'pan-flute',
            'shawm',
            'viol'
        }
    })
end
