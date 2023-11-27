import pickle
import numpy as np


def convertPklToNpz():
    with open("data/sample_weight.pkl", "rb") as f:
        data = pickle.load(f)
        np.savez("data/sample_weight", **data)


def convertNpzToNpys():
    with open("data/sample_weight.npz", "rb") as f:
        data = np.load(f)
        for key, value in data.items():
            np.save("data/weights/{}.npy".format(key), value)


if __name__ == "__main__":
    # convertPklToNpz()
    convertNpzToNpys()
