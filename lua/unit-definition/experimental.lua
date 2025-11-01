generators = {}
core = _G.core or {}

function chain(state, stack)
    local step = state['chain_step'] or 1
    local steps = {}

    local pipe
    pipe = function(fn)
        local index = #steps
        table.insert(steps, function()
            local result = fn(state['chain_arg'+index])
            if #steps > index then
                stack.push = result
                stack.target = 'chain_arg'+(index+1)
                state['chain_step']=index+1
                return
            end
            stack.pop = result
        end)
        return {
            pipe=pipe
        }
    end

    return {
        pipe=pipe,
        exec=function()
            steps[step]()
        end
    }
end

procedures = {
    unitAddBase=function(unit)
        return {
            procedure='Unit.addBase',
            args = {
                unit=unit
            }
        }
    end,
   unitAddAbilities=function(unit)
       return {
           procedure='Unit.addAbilities',
           args={
               unit=unit
           }
       }
   end,
    unitAddBackground=function(unit)
        return {
            procedure='Unit.addBackground',
            args={
                unit
            }
        }
    end,
    unitAddRace=function(unit)
        return {
            procedure='Unit.addRace',
            args={
                unit
            }
        }
    end,
    unitAddFighter_1=function(unit)
        return {
            procedure='Unit.addFighter_1',
            args={
                unit
            }
        }
    end,
    unitAddCharacterName=function(unit)
        return {
            procedure='Unit.addCharacterName',
            args={
                unit
            }
        }
    end,
    worldAddUnit=function(unit)
        return {
            procedure='World.addUnit',
            args={
                unit
            }
        }
    end
}

generators['Unit.createCharacter'] = function(args, state, ctx)
    local unit =  core.get(state, 'unit', args.unit)

    chain(state, ctx)
    .pipe(function()
        return procedures.unitAddBase(unit)
    end)
    .pipe(function(payload)
        core.merge(unit, payload)
        return procedures.unitAddAbilities(unit)
    end)
    .pipe(function(payload)
        core.merge(unit, payload)
        return procedures.unitAddBackground(unit)
    end)
    .pipe(function(payload)
        core.merge(unit, payload)
        return procedures.unitAddRace(unit)
    end)
    .pipe(function(payload)
        core.merge(unit, payload)
        return procedures.unitAddFighter_1(unit)
    end)
    .pipe(function(payload)
        core.merge(unit, payload)
        return procedures.unitAddCharacterName(unit)
    end)
    .pipe(function(payload)
        core.merge(unit, payload)
        return procedures.worldAddUnit(unit)
    end)
    .pipe(function(unitId)
        return {
            unitId=unitId
        }
    end)
    .exec()
end

