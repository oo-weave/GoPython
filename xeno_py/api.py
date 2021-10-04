# based on xeno-canto api
# see: https://www.xeno-canto.org/explore/api
import requests


class Query:
    def __init__(self, q=None, **kwargs):
        self.__get_url(q, **kwargs)

    def __get_url(self, q, **kwargs):
        url = "https://www.xeno-canto.org/api/2/recordings?query="
        url += q if q else ""
        url += "+" if q and kwargs else ""
        url += "+".join([f"{k}:{v}" for k, v in kwargs.items()])
        self.url = url

    def get(self, page=None):
        url = f"{self.url}&page={page}" if page else self.url
        return requests.get(url).json()

    def get_all(self):
        collector = []
        result = self.get()
        print(f"Processing page 1")
        collector += result['recordings']

        for page in range(2, result['numPages'] + 1):
            print(f"Processing page {page}")
            pr = self.get(page=page)
            collector += pr['recordings']
        return collector

# class API:
#     def __init__(self):
#         self.base = "https://www.xeno-canto.org/api/2"
#
#     def search(self, q=None, page=None, **kwargs):
#         url = self.get_url(q, page, **kwargs)
#         return requests.get(url).json()
#
#     def get_url(self, q=None, page=None, **kwargs):
#         url = f"{self.base}/recordings?query="
#         url += q if q else ""
#         url += "+" if q and kwargs else ""
#         url += "+".join([f"{k}:{v}" for k, v in kwargs.items()])
#         url = self.__paginate(url, page)
#         return url
#
#     @staticmethod
#     def __paginate(url, page):
#         return f"{url}&page={page}" if page else url


if __name__ == '__main__':
    api = Query()

    # Search with parameters
    rslt = api.search("Tringa", area="europe", since=10)
    print(f"Found {len(rslt['recordings'])} results with {rslt['numSpecies']} species")

    # Paginated search
    rslt = api.search(area="europe", since=10, page=2)
    print(f"Found {len(rslt['recordings'])} results on page {rslt['page']} of {rslt['numPages']}")
