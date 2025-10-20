generators = {}
core = _G.core or {}

generators['Unit.addAbilities'] = function (args, state, stack)
    local abilities = {'strength', 'constitution', 'dexterity', 'intelligence', 'charisma', 'wisdom'}
    local aidx = core.findIndex(abilities, state['ability']) + 1
    if aidx > #abilities then
        stack.pop = core.omit()
    end
    local ability = abilities[aidx]
    stack.push = {
        procedure="option.scanf",
        args={
            name = 'strength',
            type = 'int'
        },
        target = ability
    }
    state['ability'] = ability
end