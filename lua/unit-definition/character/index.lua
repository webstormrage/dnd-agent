generators = {}
core = _G.core or {}

generators['Unit.addCharacterName'] = function(args, state, stack)
    local step = state['step'] or 'name'
    local attributes = core.get(state, 'attributes', args.unit.attributes)

    local steps = {
        ['name']=function()
            stack.push = {
                procedure = "option.scanf",
                args = {
                    name = 'name',
                    type = 'string'
                }
            }
            stack.target = 'name'
            state['step'] = 'rest'
        end,
        ['rest']=function()
            local name = state['name']
            attributes['name'] = name
            stack.pop = {
                unit={
                    attributes=attributes
                }
            }
        end
    }

    steps[step]()
end