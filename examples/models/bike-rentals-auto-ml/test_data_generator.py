import pandas as pd
import argparse

def csv_to_json(input_csv, output_json):
    df = pd.read_csv(input_csv)

    # Rename the year column
    df = df.rename(columns={"yr": "year"})

    # Extract 'day' from the 'date' column
    df['day'] = df['dteday'].str.split('-').str[-1].astype(int)

    # Drop some of the columns which aren't needed
    columns_to_drop = ["instant", "dteday", "atemp", "casual", "registered", "cnt"]
    df = df.drop(columns=columns_to_drop, errors='ignore')

    # Randomly sample 20 rows from the DataFrame
    df_sample = df.sample(n=20)

    data = {
        "dataframe_split": {
            "columns": list(df_sample.columns),
            "data": df_sample.values.tolist()
        }
    }

    with open(output_json, 'w') as json_file:
        json_file.write(str(data))

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Convert CSV to JSON format.")
    parser.add_argument("input_csv", help="Path to the input CSV file.")
    parser.add_argument("output_json", help="Path to the output JSON file.")

    args = parser.parse_args()

    csv_to_json(args.input_csv, args.output_json)
