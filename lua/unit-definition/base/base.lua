function unitDefinition(attributes, inventory, options)
    -- Действия
    if attributes['actions'] == nil then
        attributes['actions'] = {}
    end

    local actions = attributes['actions']
    actions['move'] = true
    actions['attack'] = true
    actions['equip'] = true
    actions['non-lethal-attack'] = true
    actions['secondary-attack'] = true
    actions['dash'] = true
    actions['disengage'] = true
    actions['throw'] = true
    actions['dodge'] = true
    actions['help'] = true
    actions['hide'] = true
    actions['search'] = true
    actions['grapple'] = true
    actions['opportunity-attack'] = true
end