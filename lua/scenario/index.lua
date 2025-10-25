generators = {}

generators['/start'] = function(args, state, stack)
    local step = state['step'] or 'create-character'
    local steps = {
        ['create-character']=function()
            stack.push = {
                procedure='Character.create'
            }
            stack.target = 'unitId'
            state['step'] = 'spawn-character'
        end,
        ['spawn-character']=function()
            stack.push = {
                procedure='Unit.spawn',
                args={
                    unitId=state.unitId,
                    gameZoneId='frey-pastion',
                    owner='player',
                    x=2,
                    y=1
                }
            }
            state['step'] = 'fulfill'
        end,
        ['fulfill']=function()
            stack.pop = true
        end
    }
    steps[step]()
end