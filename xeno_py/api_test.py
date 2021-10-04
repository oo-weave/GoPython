import unittest
from .api import Query


class QueryTests(unittest.TestCase):

    def setUp(self) -> None:
        self.url = "https://www.xeno-canto.org/api/2"

    def test_Query_query(self):
        q = Query(q="mallard")
        self.assertEqual(f"{self.url}/recordings?query=mallard", q.url)

    def test_Query_query_kwarg(self):
        q = Query(q="pigeon", area="europe")
        self.assertEqual(f"{self.url}/recordings?query=pigeon+area:europe", q.url)

    def test_Query_query_kwargs(self):
        q = Query(q="mallard", area="europe", since=1)
        self.assertEqual(f"{self.url}/recordings?query=mallard+area:europe+since:1", q.url)

    def test_Query_kwarg(self):
        q = Query(since=1)
        self.assertEqual(f"{self.url}/recordings?query=since:1", q.url)

    def test_Query_kwargs(self):
        q = Query(area="europe", since=1)
        self.assertEqual(f"{self.url}/recordings?query=area:europe+since:1", q.url)


if __name__ == '__main__':
    unittest.main()
