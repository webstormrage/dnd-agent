function unitDefinition(attributes, equipment, options)
    -- Владения
    if attributes['proficiencies'] == nil then
        attributes['proficiencies'] = {}
    end
    attributes['proficiencies']['survival'] =  true
    attributes['proficiencies']['athletics'] = true
    attributes['proficiencies'][options['language']] = true
    attributes['proficiencies'][options['musical-skill']] = true

    -- Экипировка
    equipment['staff'] = (equipment['staff'] or 0) + 1
    equipment['hunting-trap'] = (equipment['hunting-trap'] or 0) + 1
    equipment['traveler-cloth'] = (equipment['traveler-cloth'] or 0) + 1
    equipment['coins'] = (equipment['coins'] or 0) + 1000
end

function optionsDefinition(attributes, choices)
    choices.insert({
        name='language',
        type='select',
        options= {
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
    choices.insert({
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