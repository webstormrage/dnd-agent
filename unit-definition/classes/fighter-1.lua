
function unitDefinition(attributes, equipment, options)
    -- Требования мультикласса
    if  attributes['base-class'] == nil
            and attributes['strength'] < 13
            and attributes['dexterity'] < 13 then
        return
    end

    -- Правило хитов
    if  attributes['base-class'] == nil then
        attributes['hit-points'] = 10
    else
        attributes['hit-points'] = attributes['hit-points'] + 6
    end

    -- Владения
    if attributes['proficiencies'] == nil then
        attributes['proficiencies'] = {}
    end

    attributes['proficiencies']['simple-weapon'] = true
    attributes['proficiencies']['martial-weapons'] = true
    attributes['proficiencies']['light-armor'] = true
    attributes['proficiencies']['medium-armor'] = true
    attributes['proficiencies']['heavy-armor'] = true
    attributes['proficiencies']['strength-save'] = true
    attributes['proficiencies']['constitution-save'] = true
    attributes['proficiencies'][options.skills[1]] = true
    attributes['proficiencies'][options.skills[2]] = true

    -- Бонус мастерства
    if attributes['proficiency-bonus'] == nil then
        attributes['proficiency-bonus']  = 0
    end
    attributes['proficiency-bonus'] = math.max(attributes['proficiency-bonus'], 2)


    -- Действия
    if attributes['actions'] == nil then
        attributes['actions'] = {}
    end
    attributes['actions']['second-will'] = true


    -- Особенности
    if attributes['feats'] == nil then
        attributes['feats'] = {}
    end
    attributes['feats'][options['feat']] = true

    -- Экипировка
    if attributes['base-class'] == nil then

        if options['armor'] == 'heavy-set' then
            equipment['chain-mail'] = (equipment['chain-mail'] or 0) + 1
        else
            equipment['leather-armor'] = (equipment['leather-armor'] or 0) + 1
            equipment['longbow'] = (equipment['longbow'] or 0) + 1
            equipment['arrow'] = (equipment['arrow'] or 0) + 20
        end

        local weapon1 = options['melee'][1]
        equipment[weapon1] = (equipment[weapon1] or 0) + 1
        local weapon2 = options['melee'][2]
        equipment[weapon2] = (equipment[weapon2] or 0) + 1

        if options['ranged'] == 'crossbow-set' then
            equipment['light-crossbow'] = (equipment['light-crossbow'] or 0) + 1
            equipment['bolt'] = (equipment['bolt'] or 0) + 20
        else
            equipment['handaxe'] = (equipment['handaxe'] or 0) + 1
        end

        if options['pack'] == 'dungeon-pack' then
            equipment['backpack'] = (equipment['backpack'] or 0) + 1
            equipment['crowbar'] = (equipment['crowbar'] or 0) + 1
            equipment['hammer'] = (equipment['hammer'] or 0) + 1
            equipment['piton'] = (equipment['piton'] or 0) + 10
            equipment['torch'] = (equipment['torch'] or 0) + 10
            equipment['daily-ration'] = (equipment['daily-ration'] or 0) + 10
            equipment['tinderbox'] = (equipment['tinderbox'] or 0) + 1
            equipment['waterskin'] = (equipment['waterskin'] or 0) + 1
            equipment['hempen-rope'] = (equipment['hempen-rope'] or 0) + 1
        else
            equipment['backpack'] = (equipment['backpack'] or 0) + 1
            equipment['bedroll'] = (equipment['bedroll'] or 0) + 1
            equipment['flask-of-oil'] = (equipment['flask-of-oil'] or 0) + 1
            equipment['torch'] = (equipment['torch'] or 0) + 10
            equipment['daily-ration'] = (equipment['daily-ration'] or 0) + 10
            equipment['tinderbox'] = (equipment['tinderbox'] or 0) + 1
            equipment['waterskin'] = (equipment['waterskin'] or 0) + 1
            equipment['rope'] = (equipment['rope'] or 0) + 1
        end
    end

    if attributes['base-class'] == nil then
        attributes['base-class'] = 'fighter'
    end
end

function optionsDefinition(attributes, choices)
    choices.insert({
        name='skills',
        limit=2,
        type='multi-select',
        options={
            'acrobatics',
            'animal-handling',
            'athletics',
            'history',
            'insight',
            'intimidation',
            'perception',
            'survival'
        }
    })
    choices.insert({
        name='feats',
        type='select',
        options={
            'fighting-style-archery',
            'fighting-style-defense',
            'fighting-style-great-weapon-fighting',
            'fighting-style-protection',
            'fighting-style-two-weapon-fighting'
        }
    })
    choices.insert({
        name='armor',
        type='select',
        options={
            'heavy-set',
            'light-set'
        }
    })
    choices.insert({
        name='melee',
        type='multi-select',
        limit=2,
        options={
            'battleaxe',
            'flail',
            'glaive',
            'greataxe',
            'greatsword',
            'halberd',
            'lance',
            'longsword',
            'maul',
            'morningstar',
            'pike',
            'rapier',
            'scimitar',
            'shortsword',
            'trident',
            'war-pick',
            'warhammer',
            'whip',
            'shield'
        }
    })
    choices.insert({
        name='ranged',
        type='select',
        options={
            'crossbow-set',
            'axe-set'
        }
    })
    choices.insert({
        name='pack',
        type='select',
        options={
            'explorer-pack',
            'dungeon-pack'
        }
    })
end