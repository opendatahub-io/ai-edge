import pandas as pd
import argparse

def csv_to_json(input_csv, output_json):
    df = pd.read_csv(input_csv)

    # The model expectes 8 float32 inputs, so we just need to drop one
    # of them
    columns_to_drop = ["population"]
    df = df.drop(columns=columns_to_drop)

    # Randomly sample 20 rows from the DataFrame
    df_sample = df.sample(n=20)

    data = {
        "data": df_sample.values.tolist()
    }

    with open(output_json, 'w') as json_file:
        json_file.write(str(data))

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Convert CSV to JSON format.")
    parser.add_argument("input_csv", help="Path to the input CSV file.")
    parser.add_argument("output_json", help="Path to the output JSON file.")

    args = parser.parse_args()

    csv_to_json(args.input_csv, args.output_json)
