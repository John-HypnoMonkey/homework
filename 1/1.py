import re

print 'Please enter login'

val = raw_input()
if re.match(r'([a-zA-Z])$|([a-zA-Z][a-zA-Z0-9.-]*[a-zA-Z0-9])$', val) and len(val) <= 20:
    print 'Good login'
else:
    print 'Bad login'

