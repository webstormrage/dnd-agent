handlers = {}

handlers['Scenario.start'] = function(ctx)
    -- подготовка карт уровней
    table.insert(ctx.next, {
        command='create-character'
    })
end

handlers['Character.On.create'] = function(ctx)
    table.insert(ctx.next, {
        command='spawn',
        spawn={
            unitId=ctx.characterOnCreate.unitId,
            gameZoneId='frey-pastion',
            owner='player',
            x=2,
            y=1
        }
    })
end