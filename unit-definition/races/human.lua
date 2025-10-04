function unitDefinition(attributes, inventory, options)
    -- Раса может быть только одна
    if attributes['race'] ~= nil then
        return
    end

    attributes['strength'] = attributes['strength'] + 1
    attributes['constitution'] = attributes['constitution'] + 1
    attributes['dexterity'] = attributes['dexterity'] + 1
    attributes['intelligence'] = attributes['intelligence'] + 1
    attributes['charisma'] = attributes['charisma'] + 1
    attributes['wisdom'] = attributes['wisdom'] + 1

    attributes['size'] = 'medium'
    attributes['speed'] = 30

    if attributes['proficiencies'] == nil then
        attributes['proficiencies'] = {}
    end

    attributes['proficiencies']['language-common'] = true
    attributes['proficiencies'][options['second-language']] = true

    attributes['race'] = 'human'
end

function optionsDefinition(attributes, choices)
    table.insert(choices, {
        name = 'second-language',
        type = 'select',
        options = {
            'language-abyssal',
            'language-celestial',
            'language-draconic',
            'language-deep-speech',
            'language-infernal',
            'language-primordial',
            'language-sylvan',
            'language-undercommon',
            'language-common',
            'language-dwarvish',
            'language-elvish',
            'language-giant',
            'language-gnomish',
            'language-goblin',
            'language-halfling',
            'language-orc'
        }
    })
end