# All material is licensed under the Apache License Version 2.0, January 2004
# http://www.apache.org/licenses/LICENSE-2.0

# python 

# Sample program to compare parsing a clean CSV in python to
# parsing a clean CSV in Go. Don't worry about the details at this stage.
import pandas as pd

# Define column names.
cols = [
        'integercolumn',
        'stringcolumn'
        ]

# Read in the CSV with pandas.
data = pd.read_csv('../data/example_clean.csv', names=cols)

# Print out the maximum value in the integer column.
print(data['integercolumn'].max())
