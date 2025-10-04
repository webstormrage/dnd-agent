
function unitDefinition(attributes, inventory, options)
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
    attributes['proficiencies'][options['primary-skill']] = true
    attributes['proficiencies'][options['secondary-skill']] = true

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
            inventory['chain-mail'] = (inventory['chain-mail'] or 0) + 1
        else
            inventory['leather-armor'] = (inventory['leather-armor'] or 0) + 1
            inventory['longbow'] = (inventory['longbow'] or 0) + 1
            inventory['arrow'] = (inventory['arrow'] or 0) + 20
        end

        local weapon1 = options['primary-melee']
        inventory[weapon1] = (inventory[weapon1] or 0) + 1
        local weapon2 = options['secondary-melee']
        inventory[weapon2] = (inventory[weapon2] or 0) + 1

        if options['ranged'] == 'crossbow-set' then
            inventory['light-crossbow'] = (inventory['light-crossbow'] or 0) + 1
            inventory['bolt'] = (inventory['bolt'] or 0) + 20
        else
            inventory['handaxe'] = (inventory['handaxe'] or 0) + 1
        end

        if options['pack'] == 'dungeon-pack' then
            inventory['backpack'] = (inventory['backpack'] or 0) + 1
            inventory['crowbar'] = (inventory['crowbar'] or 0) + 1
            inventory['hammer'] = (inventory['hammer'] or 0) + 1
            inventory['piton'] = (inventory['piton'] or 0) + 10
            inventory['torch'] = (inventory['torch'] or 0) + 10
            inventory['daily-ration'] = (inventory['daily-ration'] or 0) + 10
            inventory['tinderbox'] = (inventory['tinderbox'] or 0) + 1
            inventory['waterskin'] = (inventory['waterskin'] or 0) + 1
            inventory['hempen-rope'] = (inventory['hempen-rope'] or 0) + 1
        else
            inventory['backpack'] = (inventory['backpack'] or 0) + 1
            inventory['bedroll'] = (inventory['bedroll'] or 0) + 1
            inventory['flask-of-oil'] = (inventory['flask-of-oil'] or 0) + 1
            inventory['torch'] = (inventory['torch'] or 0) + 10
            inventory['daily-ration'] = (inventory['daily-ration'] or 0) + 10
            inventory['tinderbox'] = (inventory['tinderbox'] or 0) + 1
            inventory['waterskin'] = (inventory['waterskin'] or 0) + 1
            inventory['rope'] = (inventory['rope'] or 0) + 1
        end
    end

    if attributes['base-class'] == nil then
        attributes['base-class'] = 'fighter'
    end
end

function optionsDefinition(attributes, choices)
    table.insert(choices, {
        name = 'primary-skill',
        type = 'select',
        options = {
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

    table.insert(choices, {
        name = 'secondary-skill',
        type = 'select',
        options = {
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

    table.insert(choices, {
        name = 'feats',
        type = 'select',
        options = {
            'fighting-style-archery',
            'fighting-style-defense',
            'fighting-style-great-weapon-fighting',
            'fighting-style-protection',
            'fighting-style-two-weapon-fighting'
        }
    })

    table.insert(choices, {
        name = 'armor',
        type = 'select',
        options = {
            'heavy-set',
            'light-set'
        }
    })

    table.insert(choices, {
        name = 'primary-melee',
        type = 'select',
        options = {
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

    table.insert(choices, {
        name = 'secondary-melee',
        type = 'select',
        options = {
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

    table.insert(choices, {
        name = 'ranged',
        type = 'select',
        options = {
            'crossbow-set',
            'axe-set'
        }
    })

    table.insert(choices, {
        name = 'pack',
        type = 'select',
        options = {
            'explorer-pack',
            'dungeon-pack'
        }
    })
end