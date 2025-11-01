generators = {}
core = _G.core or {}

generators['Unit.addRace'] = function(args, state, stack)
    local step = state['step'] or 'language'
    local attributes = core.get(state, 'attributes', args.unit.attributes)
    local proficiencies = core.get(state, 'proficiencies', attributes['proficiencies'])
    attributes['proficiencies'] = proficiencies

    if attributes['race'] ~= nil  then
        stack.pop = {
            unit = {
                attributes=attributes
            }
        }
        return
    end

    local steps = {
        ['language']=function()
            local options = core.filterBy(core.LANGS, proficiencies)
            if #options == 0 then
                state['step'] = 'rest'
            end
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'language',
                    type = 'select',
                    options = core.filterBy(core.LANGS, proficiencies)
                }
            }
            stack.target = 'language'
            state['step'] = 'rest'
        end,
        ['rest']=function()
            local lang = state['language']
            if lang ~= nil then
                proficiencies[lang] = true
            end

            attributes['strength'] = (attributes['strength'] or 0) + 1
            attributes['constitution'] = (attributes['constitution'] or 0) + 1
            attributes['dexterity'] = (attributes['dexterity'] or 0) + 1
            attributes['intelligence'] = (attributes['intelligence'] or 0) + 1
            attributes['charisma'] = (attributes['charisma'] or 0) + 1
            attributes['wisdom'] = (attributes['wisdom'] or 0) + 1
            attributes['size'] = 'medium'
            attributes['speed'] = 30

            proficiencies['language/common'] = true

            attributes['race'] = 'human'
            stack.pop = {
                unit={
                    attributes=attributes
                }
            }
        end
    }

    steps[step]()
end