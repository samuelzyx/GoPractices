The following input contains a list of doctors and labels that identify their field and other characteristics, like if they are taking new patients, or if they board certified.

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

Write a function that returns all the oncologists that are board certified ordered alphabetically, like this:

['Adrien', 'Emma']