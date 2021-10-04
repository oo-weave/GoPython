from xeno_py import Query


if __name__ == '__main__':

    # search all of europe for last 20 days
    r = Query(area="europe", since=20).get_all()
    print(f"Results found: {len(r)}")
