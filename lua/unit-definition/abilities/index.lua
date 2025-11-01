generators = {}
core = _G.core or {}

generators['Unit.addAbilities'] = function (args, state, stack)
    local abilities = {'strength', 'constitution', 'dexterity', 'intelligence', 'charisma', 'wisdom'}
    local attributes = core.get(state, 'attributes', args.unit.attributes)
    local step = state['step']
    if step ~= nil then
        attributes[step] = state[step]
    end
    local aidx = core.findIndex(abilities, step) + 1
    if aidx > #abilities then
        stack.pop = {
            unit={
                attributes=abilities
            }
        }
    end
    local ability = abilities[aidx]
    stack.push = {
        procedure="option.scanf",
        args={
            name = 'strength',
            type = 'int'
        }
    }
    stack.target = ability
    state['step'] = ability
end