# All material is licensed under the Apache License Version 2.0, January 2004
# http://www.apache.org/licenses/LICENSE-2.0

# python 

# Sample program to illustrate a breakdown in integrity when parsing
# a CSV file with python.
import pandas as pd

# Define column names.
cols = [
        'integercolumn',
        'stringcolumn'
        ]

# Read in the CSV with pandas.
data = pd.read_csv('../data/example_messy.csv', names=cols)

# Print out the maximum value in the integer column.
print(data['integercolumn'].max())
