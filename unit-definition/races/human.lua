function unitDefinition(attributes, equipment, options)
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
    choices.insert({
        name='second-language',
        type='select',
        options= {
            'abyssal',
            'celestial',
            'draconic',
            'deep-speech',
            'infernal',
            'primordial',
            'sylvan',
            'undercommon',
            'common',
            'dwarvish',
            'elvish',
            'giant',
            'gnomish',
            'goblin',
            'halfling',
            'orc'
        }
    })
end