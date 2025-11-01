generators = {}
core = _G.core or {}

generators['Unit.createCharacter'] = function(args, state, stack)
    local step = state['step'] or 'base'
    local unit =  core.get(state, 'unit', args.unit)

    local steps = {
        ['base']=function()
            stack.push = {
                procedure='Unit.addBase',
                args = {
                    unit=unit
                }
            }
            stack.target = 'payload'
            state['step'] = 'abilities'
        end,
        ['abilities']=function()
            stack.push = {
                procedure='Unit.addAbilities',
                args = {
                    unit=unit
                }
            }
            stack.target = 'payload'
            state['step'] = 'background'
        end,
        ['background']=function()
            core.merge(unit, state['payload'])
            stack.push = {
                procedure='Unit.addBackground',
                args={
                    unit
                }
            }
            stack.target = 'payload'
            state['step'] = 'race'
        end,
        ['race']=function()
            core.merge(unit, state['payload'])
            stack.push = {
                procedure='Unit.addRace',
                args={
                    unit
                }
            }
            stack.target = 'payload'
            state['step'] = 'class'
        end,
        ['class']=function()
            core.merge(unit, state['payload'])
            stack.push = {
                procedure='Unit.addFighter_1',
                args={
                    unit
                }
            }
            stack.target = 'payload'
            state['step'] = 'character'
        end,
        ['character']=function()
            core.merge(unit, state['payload'])
            stack.push = {
                procedure='Unit.addCharacter',
                args={
                    unit
                }
            }
            stack.target = 'payload'
            state['step'] = 'save'
        end,
        ['save']=function()
            core.merge(unit, state['payload'])
            stack.push = {
                procedure='World.addUnit',
                args={
                    unit
                }
            }
            stack.target = 'unitId'
            state['step'] = 'rest'
        end,
        ['rest']=function()
            local unitId = state['unitId']
            stack.pop = {
                unitId=unitId
            }
        end
    }
    steps[step]()
end