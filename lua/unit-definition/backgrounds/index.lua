generators = {}
core = _G.core or {}

generators['Unit.addBackground'] = function (args, state, stack)
    local step = state['step'] or 'language'
    local attributes = core.get(state, 'attributes', {})
    local inventory = core.get(state, 'inventory', {})
    local steps = {
        ['language'] = function()
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'language',
                    type = 'select',
                    options = core.filterBy(core.LANGS, args.attributes['proficiency'])
                }
            }
            stack.target = 'language'
            state['step'] = 'musical-skill'
        end,
        ['musical-skill'] = function()
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'musical-skill',
                    type = 'select',
                    options = core.filterBy(core.MUSIC, args.attributes['proficiency'])
                },
            }
            stack.target = 'musical-skill'
            state['step'] = 'rest'
        end,
        ['rest'] = function()
            local lang = state['language']
            if lang ~= nil then
                attributes['proficiencies'][lang] = true
            end

            local music = state['musical-skill']
            if music ~= nil then
                attributes['proficiencies'][music] = true
            end

            attributes['proficiencies']['survival'] =  true
            attributes['proficiencies']['athletics'] = true

            inventory['staff'] = (inventory['staff'] or 0) + 1
            inventory['hunting-trap'] = (inventory['hunting-trap'] or 0) + 1
            inventory['traveler-cloth'] = (inventory['traveler-cloth'] or 0) + 1
            inventory['coins'] = (inventory['coins'] or 0) + 1000

            stack.pop = {
                attributes=attributes,
                inventory=inventory
            }
        end
    }
    steps[step]()
end