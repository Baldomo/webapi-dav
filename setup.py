from setuptools import setup

setup(
    name='webapi-builder',
    version='0.1',
    py_modules=['builder'],
    include_package_data=True,
    install_requires=['click'],
    entry_points='''
        [console_scripts]
        builder=builder:cli
    '''
)
