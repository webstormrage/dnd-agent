handlers = {}

handlers['/start'] = function(ctx)
    table.insert(ctx.next, {
        command='Character.create'
    })
end

-- TODO listeners['Character.create']
handlers['Character.On.create'] = function(ctx)
    table.insert(ctx.next, {
        command='Unit.spawn',
        spawn={
            unitId=ctx.characterOnCreate.unitId,
            gameZoneId='frey-pastion',
            owner='player',
            x=2,
            y=1
        }
    })
end