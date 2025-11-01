generators = {}
core = _G.core or {}

generators['Unit.addBase'] = function(args, state, stack)
    stack.pop = {
       unit={
           attributes={
               actions={
                   ['move']=true,
                   ['attack']=true,
                   ['equip']=true,
                   ['non-lethal-attack']=true,
                   ['secondary-attack']=true,
                   ['dash']=true,
                   ['disengage']=true,
                   ['throw']=true,
                   ['dodge']=true,
                   ['hide']=true,
                   ['search']=true,
                   ['grapple']=true,
                   ['opportunity-attack']=true
               }
           }
       }
    }
end