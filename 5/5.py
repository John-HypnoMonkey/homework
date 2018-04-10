keys = ['one', 'two', 'three', 'four']
vals = [1, 2, 3, 4, 5, 6, 7]
vals2 = [1, 2]


def createDict(keys, vals):

    result_dict = {}
    i = 0
    for key in keys:
        if len(vals) <= i:
            val = None
        else:
            val = vals[i]
        result_dict[str(key)] = val
        i += 1
        result_dict.clear
    return result_dict


dict1 = createDict(keys, vals)
dict2 = createDict(keys, vals2)
print(dict1)
print(dict2)
