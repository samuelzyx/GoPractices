This time, the input will change, you'll still get the same list of doctors

[
    ['name', 'labels'],
    ['Matt', 'board certified, primary care, male, takes_new_patients'],
    ['Belda', 'board certified, internal medicine, female'],
    ['Wyatt', 'primary care, male, takes_new_patients'],
    ['Emma', 'board certified, oncology'],
    ['Aaron', 'sanctioned, primary care'],
    ['Josh', 'board certified, internal medicine, takes_new_patients'],
    ['Adrien', 'oncology, board certified, takes_new_patients'],
    ['Andy', 'internal medicine, male, sanctioned']
]

However, you'll also receive a second parameter with thhe search parameters, an array of strings that contains the labels you want to search by:

['primary care', 'not sanctioned']

Generalize the function to be able to take arbitrary combinations of requirements.

Excluding paramaters are labeled with the word "not", meaning if you get a "not male" search parameter, you must EXCLUDE all doctors with label "male"

Example:

['primary care', 'not sanctioned']

Will return only primary care doctors who are NOT sactioned

['Matt', 'Wyatt']