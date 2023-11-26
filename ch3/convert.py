import pickle
import numpy as np


def convert():
    with open("data/sample_weight.pkl", "rb") as f:
        data = pickle.load(f)
        np.savez("data/sample_weight", **data)


if __name__ == "__main__":
    convert()
