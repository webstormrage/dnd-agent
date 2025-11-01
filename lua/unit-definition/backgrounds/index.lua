generators = {}
core = _G.core or {}

generators['Unit.addBackground'] = function (args, state, stack)
    local step = state['step'] or 'language'
    local attributes = core.get(state, 'attributes', args.unit.attributes)
    local inventory = core.get(state, 'inventory', args.unit.inventory)
    local proficiencies = core.get(state, 'proficiencies', attributes['proficiencies'])
    attributes['proficiencies'] = proficiencies

    local steps = {
        ['language'] = function()
            local options =  core.filterBy(core.LANGS, proficiencies)
            if #options == 0 then
                state['step'] = 'musical-skill'
                return
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'language',
                    type = 'select',
                    options = options
                }
            }
            stack.target = 'language'
            state['step'] = 'musical-skill'
        end,
        ['musical-skill'] = function()
            local lang = state['language']
            if lang ~= nil then
                proficiencies[lang] = true
            end
            local options =  core.filterBy(core.MUSIC, proficiencies)
            if #options == 0 then
                state['step'] = 'rest'
                return
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'musical-skill',
                    type = 'select',
                    options = options
                },
            }
            stack.target = 'musical-skill'
            state['step'] = 'rest'
        end,
        ['rest'] = function()
            local music = state['musical-skill']
            if music ~= nil then
                proficiencies[music] = true
            end

            proficiencies['survival'] =  true
            proficiencies['athletics'] = true

            inventory['staff'] = (inventory['staff'] or 0) + 1
            inventory['hunting-trap'] = (inventory['hunting-trap'] or 0) + 1
            inventory['traveler-cloth'] = (inventory['traveler-cloth'] or 0) + 1
            inventory['coins'] = (inventory['coins'] or 0) + 1000

            stack.pop = {
                unit={
                    attributes=attributes,
                    inventory=inventory
                }
            }
        end
    }
    steps[step]()
end