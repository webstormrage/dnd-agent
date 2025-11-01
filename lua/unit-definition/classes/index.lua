generators = {}
core = _G.core or {}

generators['Unit.addFighter_1'] = function(args, state, stack)
    local step = state['step'] or 'primary-skill'
    local attributes = core.get(state, 'attributes', args.unit.attributes)
    local inventory = core.get(state, 'inventory', args.unit.inventory)
    local proficiencies = core.get(attributes, 'proficiencies', attributes['proficiencies'])
    attributes['proficiencies'] = proficiencies
    local feats = core.get(attributes, 'feats', attributes['feats'])
    attributes['feats'] = feats

    -- Требования мультикласса
    if  attributes['base-class'] == nil
            and attributes['strength'] < 13
            and attributes['dexterity'] < 13 then
        stack.pop = {
            unit={
                attributes,
                inventory
            }
        }
        return
    end

    local steps = {
        ['primary-skill']=function()
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'primary-skill',
                    type = 'select',
                    options = core.filterBy({
                        'acrobatics',
                        'animal-handling',
                        'athletics',
                        'history',
                        'insight',
                        'intimidation',
                        'perception',
                        'survival'
                    }, proficiencies)
                }
            }
            stack.target = 'primary-skill'
            state['step'] = 'secondary-skill'
        end,
        ['secondary-skill']=function()
            local primarySkill = state['primary-skill']
            if primarySkill ~= nil then
                proficiencies[primarySkill] = true
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'secondary-skill',
                    type = 'select',
                    options = core.filterBy({
                        'acrobatics',
                        'animal-handling',
                        'athletics',
                        'history',
                        'insight',
                        'intimidation',
                        'perception',
                        'survival'
                    }, proficiencies)
                }
            }
            stack.target = 'secondary-skill'
            state['step'] = 'feat'
        end,
        ['feat']=function()
            local secondarySkill = state['secondary-skill']
            if secondarySkill ~= nil then
                proficiencies[secondarySkill] = true
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'feat',
                    type = 'select',
                    options = core.filterBy(core.FIGHTING_STYLES, attributes)
                }
            }
            stack.target = 'feat'
            state['step'] = 'armor'
        end,
        ['armor']=function()
            local feat = state['feat']
            if feat ~= nil then
                feats[feat] = true
            end
            if attributes['base-class'] ~= nil then
                state['step'] = 'rest'
                return
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'armor',
                    type = 'select',
                    options = core.merge({
                        'heavy-set',
                        'light-set'
                    })
                }
            }
            stack.target = 'armor'
            state['step'] = 'primary-melee'
        end,
        ['primary-melee']=function()
            local armorSet = state['armor']
            if armorSet == 'heavy-set' then
                inventory['chain-mail'] = (inventory['chain-mail'] or 0) + 1
            else
                inventory['armor/leather'] = (inventory['armor/leather'] or 0) + 1
                inventory['longbow'] = (inventory['longbow'] or 0) + 1
                inventory['arrow'] = (inventory['arrow'] or 0) + 20
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'primary-melee',
                    type = 'select',
                    options = core.merge({ 'shield' }, core.MARTIAL_MEELE_WEAPONS)
                }
            }
            stack.target = 'primary-melee'
            state['step'] = 'secondary-melee'
        end,
        ['secondary-melee']=function()
            local melee = state['primary-melee']
            inventory[melee] = (inventory[melee] or 0) + 1
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'secondary-melee',
                    type = 'select',
                    options = core.merge({ 'shield' }, core.MARTIAL_MEELE_WEAPONS)
                }
            }
            stack.target = 'secondary-melee'
            state['step'] = 'ranged'
        end,
        ['ranged']=function()
            local melee = state['secondary-melee']
            inventory[melee] = (inventory[melee] or 0) + 1
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'ranged',
                    type = 'select',
                    options = core.merge({ 'crossbow-set' ,'axe-set' })

                }
            }
            stack.target = 'ranged'
            state['step'] = 'pack'
        end,
        ['pack']=function()
            local ranged = state['ranged']
            if ranged == 'crossbow-set' then
                inventory['light-crossbow'] = (inventory['light-crossbow'] or 0) + 1
                inventory['bolt'] = (inventory['bolt'] or 0) + 20
            else
                inventory['handaxe'] = (inventory['handaxe'] or 0) + 1
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'pack',
                    type = 'select',
                    options = core.merge({ 'explorer-pack' ,'dungeon-pack' })
                }
            }
            stack.target = 'pack'
            state['step'] = 'rest'
        end,
        ['rest']=function()
            local pack = state['pack']
            if pack == 'dungeon-pack' then
                inventory['backpack'] = (inventory['backpack'] or 0) + 1
                inventory['crowbar'] = (inventory['crowbar'] or 0) + 1
                inventory['hammer'] = (inventory['hammer'] or 0) + 1
                inventory['piton'] = (inventory['piton'] or 0) + 10
                inventory['torch'] = (inventory['torch'] or 0) + 10
                inventory['daily-ration'] = (inventory['daily-ration'] or 0) + 10
                inventory['tinderbox'] = (inventory['tinderbox'] or 0) + 1
                inventory['waterskin'] = (inventory['waterskin'] or 0) + 1
                inventory['hempen-rope'] = (inventory['hempen-rope'] or 0) + 1
            elseif pack == 'explorer-pack' then
                inventory['backpack'] = (inventory['backpack'] or 0) + 1
                inventory['bedroll'] = (inventory['bedroll'] or 0) + 1
                inventory['flask-of-oil'] = (inventory['flask-of-oil'] or 0) + 1
                inventory['torch'] = (inventory['torch'] or 0) + 10
                inventory['daily-ration'] = (inventory['daily-ration'] or 0) + 10
                inventory['tinderbox'] = (inventory['tinderbox'] or 0) + 1
                inventory['waterskin'] = (inventory['waterskin'] or 0) + 1
                inventory['rope'] = (inventory['rope'] or 0) + 1
            end
            attributes['base-class'] = 'fighter'
            stack.pop = {
                unit = {
                    attributes=attributes,
                    inventory=inventory
                }
            }
        end
    }

    steps[step]()
end